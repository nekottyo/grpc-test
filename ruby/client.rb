require 'grpc'
require_relative 'time_services_pb'

def main
  stub = Grpctest::Time::V1::TimeService::Stub.new('localhost:50051', :this_channel_is_insecure)
  request = Grpctest::Time::V1::GetCurrentTimeRequest.new
  response = stub.get_current_time(request)
  p response.date
end

main
