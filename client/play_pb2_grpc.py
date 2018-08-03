# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import play_pb2 as play__pb2


class TurnStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.SetMove = channel.unary_unary(
        '/play.Turn/SetMove',
        request_serializer=play__pb2.Step.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.GetMove = channel.unary_unary(
        '/play.Turn/GetMove',
        request_serializer=play__pb2.Player.SerializeToString,
        response_deserializer=play__pb2.Step.FromString,
        )
    self.UpdateNext = channel.unary_unary(
        '/play.Turn/UpdateNext',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.IsNextPlayer = channel.unary_unary(
        '/play.Turn/IsNextPlayer',
        request_serializer=play__pb2.Player.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.SetPlayer = channel.unary_unary(
        '/play.Turn/SetPlayer',
        request_serializer=play__pb2.Player.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.GetAIPlayer = channel.unary_unary(
        '/play.Turn/GetAIPlayer',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.Player.FromString,
        )
    self.HasChosen = channel.unary_unary(
        '/play.Turn/HasChosen',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.HasMoved = channel.unary_unary(
        '/play.Turn/HasMoved',
        request_serializer=play__pb2.Player.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.SetResumed = channel.unary_unary(
        '/play.Turn/SetResumed',
        request_serializer=play__pb2.Resumed.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.GetResumed = channel.unary_unary(
        '/play.Turn/GetResumed',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.Resumed.FromString,
        )
    self.NewRoom = channel.unary_unary(
        '/play.Turn/NewRoom',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.GetID = channel.unary_unary(
        '/play.Turn/GetID',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.SetExit = channel.unary_unary(
        '/play.Turn/SetExit',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.CheckExit = channel.unary_unary(
        '/play.Turn/CheckExit',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.SetUserID = channel.unary_unary(
        '/play.Turn/SetUserID',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )
    self.GetUserID = channel.unary_unary(
        '/play.Turn/GetUserID',
        request_serializer=play__pb2.State.SerializeToString,
        response_deserializer=play__pb2.State.FromString,
        )


class TurnServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def SetMove(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetMove(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateNext(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def IsNextPlayer(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetPlayer(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetAIPlayer(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def HasChosen(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def HasMoved(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetResumed(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetResumed(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def NewRoom(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetID(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetExit(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CheckExit(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetUserID(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetUserID(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_TurnServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SetMove': grpc.unary_unary_rpc_method_handler(
          servicer.SetMove,
          request_deserializer=play__pb2.Step.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'GetMove': grpc.unary_unary_rpc_method_handler(
          servicer.GetMove,
          request_deserializer=play__pb2.Player.FromString,
          response_serializer=play__pb2.Step.SerializeToString,
      ),
      'UpdateNext': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateNext,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'IsNextPlayer': grpc.unary_unary_rpc_method_handler(
          servicer.IsNextPlayer,
          request_deserializer=play__pb2.Player.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'SetPlayer': grpc.unary_unary_rpc_method_handler(
          servicer.SetPlayer,
          request_deserializer=play__pb2.Player.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'GetAIPlayer': grpc.unary_unary_rpc_method_handler(
          servicer.GetAIPlayer,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.Player.SerializeToString,
      ),
      'HasChosen': grpc.unary_unary_rpc_method_handler(
          servicer.HasChosen,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'HasMoved': grpc.unary_unary_rpc_method_handler(
          servicer.HasMoved,
          request_deserializer=play__pb2.Player.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'SetResumed': grpc.unary_unary_rpc_method_handler(
          servicer.SetResumed,
          request_deserializer=play__pb2.Resumed.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'GetResumed': grpc.unary_unary_rpc_method_handler(
          servicer.GetResumed,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.Resumed.SerializeToString,
      ),
      'NewRoom': grpc.unary_unary_rpc_method_handler(
          servicer.NewRoom,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'GetID': grpc.unary_unary_rpc_method_handler(
          servicer.GetID,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'SetExit': grpc.unary_unary_rpc_method_handler(
          servicer.SetExit,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'CheckExit': grpc.unary_unary_rpc_method_handler(
          servicer.CheckExit,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'SetUserID': grpc.unary_unary_rpc_method_handler(
          servicer.SetUserID,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
      'GetUserID': grpc.unary_unary_rpc_method_handler(
          servicer.GetUserID,
          request_deserializer=play__pb2.State.FromString,
          response_serializer=play__pb2.State.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'play.Turn', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
