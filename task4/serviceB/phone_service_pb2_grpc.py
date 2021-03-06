# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import phone_service_pb2 as phone__service__pb2


class PhoneServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CheckPhone = channel.unary_unary(
                '/PhoneService/CheckPhone',
                request_serializer=phone__service__pb2.CheckPhoneRequest.SerializeToString,
                response_deserializer=phone__service__pb2.CheckPhoneResponse.FromString,
                )


class PhoneServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CheckPhone(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PhoneServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CheckPhone': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckPhone,
                    request_deserializer=phone__service__pb2.CheckPhoneRequest.FromString,
                    response_serializer=phone__service__pb2.CheckPhoneResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'PhoneService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PhoneService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CheckPhone(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/PhoneService/CheckPhone',
            phone__service__pb2.CheckPhoneRequest.SerializeToString,
            phone__service__pb2.CheckPhoneResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
