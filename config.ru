require 'sinatra'

class App < Sinatra::Base

  helpers do
    def halter(code)
      halt code.to_i
    end
  end

  get "/:code" do
    halter params[:code]
  end

  post "/:code" do
    halter params[:code]
  end

  put "/:code" do
    halter params[:code]
  end

  delete "/:code" do
    halter params[:code]
  end

  patch "/:code" do
    halter params[:code]
  end

end

run App
