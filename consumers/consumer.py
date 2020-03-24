import copy
import logging
from abc import ABC, abstractmethod
from pathlib import Path

from api.proto import engine_pb2
from third_party.python.google.protobuf.message import DecodeError

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


# A base consumer class that all implementations of consumers should implement
class Consumer(ABC):

    @abstractmethod
    def __init__(self, config: dict):
        try:
            self.pvc_location = config['pvc_location']
        except(KeyError):
            logger.error('PVC location not provided')
            raise

        logger.info('Instantiated Consumer class with results at ' + self.pvc_location)

    @abstractmethod
    def load_results(self):
        """
        Load a set of LaunchToolResponse protobufs into a list for processing
        """

        scan_results = engine_pb2.EnrichedLaunchToolResponse()
        collected_results = self.load_files(scan_results, self.pvc_location)

        return collected_results

    @abstractmethod
    def send_results(self, collected_results: list):
        """
        Implementations should send results to their platforms, the implementation
        below merely logs the results
        """

        for scan in collected_results:
            raw_scan = scan.original_results
            scan_time = raw_scan.scan_info.scan_start_time.ToJsonString()
            logger.info('Scan: ' + raw_scan.tool_name + ' run at ' + scan_time)
            for issue in raw_scan.issues:
                logger.info('Issue: ' + str(issue))




    def load_files(self, protobuf, location):
        """Given a protobuf object and a filesystem location, attempts to load all *.pb
        files found in directories underneath the location into the protobuf object

        :param protobuf: object expected to be found in the location
        :param location: directory where protobuf objects are stored

        :returns array of protobuf objects of the given type which were found at the location
        :raise SyntaxError: If there are no .pb files found in location
        """

        logger.info('Searching for scan results')
        collected_files = []

        for filename in  Path(location).glob('**/*.pb'):
            logger.info("Found file %s" % filename)
            with open(filename, "rb") as f:
                try:
                    protobuf.ParseFromString(f.read())
                    collected_files.append(copy.deepcopy(protobuf))

                except DecodeError as e:
                    logger.warning('Unable to parse file %s skipping because of: %s '%(filename,str(e)))
                    # Note: here skipping is important,
                    #  the results dir might have all sorts of protobuf messages that don't
                    #  match the type provided
        if len(collected_files) == 0:
            raise SyntaxError('No valid results were found in the provided directory')

        return collected_files
