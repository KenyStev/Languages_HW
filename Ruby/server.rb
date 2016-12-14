require 'sinatra'
require 'json'
require_relative 'gp'
require_relative 'services/services'

$SORTPATH = "resources/mergesort/"
$BITCODEPATH = "resources/bitcode/"

get '/' do
  'Hello world!'
end

post '/api/sort' do
  upload(params,$SORTPATH)
  send_file SortEmails(params[:file][:filename])
end

post '/api/bitcode' do
  upload(params,$BITCODEPATH)
  send_file HideMessage(params[:file][:filename],params[:message])
end

post '/api/bitcode/seek' do
  upload(params,$BITCODEPATH)
  send_file SeekMessage(params[:file][:filename])
end

post '/api/kruskal' do
  request.body.rewind
  payload = JSON.parse(request.body.read)
  Kruskal(payload)
end