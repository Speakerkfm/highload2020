import grpc
from concurrent import futures
import time

import phone_service_pb2_grpc as pb2_grpc
import phone_service_pb2 as pb2


class PhoneService(pb2_grpc.PhoneServiceServicer):

    def __init__(self, *args, **kwargs):
        pass

    def CheckPhone(self, request, context):
        print(request)
        exists = request.phone == "+79125947072"

        return pb2.CheckPhoneResponse(exists=exists)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_PhoneServiceServicer_to_server(PhoneService(), server)
    server.add_insecure_port('0.0.0.0:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
