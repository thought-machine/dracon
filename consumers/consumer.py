import logging

from abc import ABC, abstractmethod

from gen import engine_pb2
from utils.file_utils import load_files

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


# A base consumer class that all implementations of consumers should implement
class Consumer(ABC):

    @abstractmethod
    def __init__(self, config: dict):
        try:
            self.pb_location = config['pb_location']
        except(KeyError):
            logger.error('PVC location not provided')
            raise

        logger.info('Instantiated Consumer class with results at ' + self.pb_location)

    @abstractmethod
    def load_results(self):
        """
        Load a set of LaunchToolResponse protobufs into a list for processing
        """

        scan_results = engine_pb2.EnrichedLaunchToolResponse()
        collected_results = load_files(scan_results, self.pb_location)

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
