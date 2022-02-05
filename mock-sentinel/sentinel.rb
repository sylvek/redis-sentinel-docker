require 'socket'

server = TCPServer.new 26379

def _answer(client, command)
  puts ">>> #{command}"
  File.readlines(command).each do |line|
    client.puts line
    puts "<<< #{line}"
  end
end

loop do
  Thread.start(server.accept) do |client|
    loop do
      first_line = client.gets
      if first_line
        first_line = first_line.delete("\r\n").downcase
        if first_line.start_with? '*'
          number_of_args = first_line[1..-1].to_i
          args = Array.new
          for arg in 0..number_of_args-1 do
            _i_dont_care = client.gets
            args << client.gets.delete("\r\n").downcase
          end
          _answer(client, args.join('-'))
        else
          _answer(client, first_line.gsub(' ', '-'))
        end
      end
    end
  end
end