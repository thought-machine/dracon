import argparse
import sys
import json

from api.proto import engine_pb2
from consumers.consumer import Consumer
from third_party.python import requests

import logging

logger = logging.getLogger(__name__)

class SimpleConsumer(Consumer):

    def __init__(self, config: dict):
        logger.info("Starting Consumer")
        self.pvc_location = config.pvc_location
        logger.info("Reading from %s" % self.pvc_location)

        if (self.pvc_location is None):
            raise AttributeError("PVC claim location is missing")

    def load_results(self) -> (list, bool):
        try:
            return self._load_enriched_results(), False
        except SyntaxError:
            return self._load_plain_results(), True

    def _load_plain_results(self):
        scan_results = engine_pb2.LaunchToolResponse()
        return self.load_files(scan_results, self.pvc_location)

    def _load_enriched_results(self):
        """Load a set of LaunchToolResponse protobufs into a list for processing"""
        return super().load_results()

    def print_data(self, fact, N):
        """Prints a simple message to stdout"""
        msg = "Here's a random fact, " + fact + " You have received " + \
              "a random fact because you have " + str(N) + " results.\n"
        print(msg)
        return msg

    def send_results(self, collected_results: list, raw: bool):
        """
        In this example, we perform a request to a url and extract a random 'fact' from the JSON response
        And then send it to print_data() for printing.

        :param collected_results: list of LaunchToolResponse protobufs
        """
        try:
            resp = requests.get(url='https://uselessfacts.jsph.pl/random.json')
            data = resp.json()
            fact = data.get('text')
            if fact is None:
                logger.error("The field 'text' is missing from the specified json response")
                sys.exit(-1)
        except requests.exceptions.RequestException as e:
            logger.error('Error while performing the http request: {}'.format(str(e)))
            sys.exit(-1)

        return self.print_data(fact=fact, N=len(collected_results))


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        '--pvc_location', help='The location of the scan results')
    parser.add_argument(
        '--raw', help='if it should process raw or enriched results', action="store_true")
    args = parser.parse_args()
    ec = SimpleConsumer(args)
    try:
        logger.info('Loading results from %s' % str(ec.pvc_location))
        # collected_results, raw = ec.load_results()
        collected_results = list()
        raw = False
        logger.info("gathered %s results", len(collected_results))
        logger.info("Reading raw: %s ", raw)
    except SyntaxError as e:
        logger.error('Unable to load results from %s: ', str(e))
        sys.exit(-1)

    ec.send_results(collected_results, args.raw)
    logger.info('Done!')


if __name__ == '__main__':
    logger.info("Simple Consumer running")
    main()