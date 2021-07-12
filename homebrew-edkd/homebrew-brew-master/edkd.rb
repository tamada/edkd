VERSION="1.0.0"

class Edkd < Formula
  desc "Calculate the distance from the latitude(ed) and longitude(kd) of two points"
  homepage "https://github.com/Ykatsuy/edkd"
  url "https://github.com/tamada/edkd/releases/download/v#{VERSION}/edkd-#{VERSION}_darwin_amd64.tar.gz"
  version VERSION
  license "MIT License"

  option "without-completions", "Disable bash completions"
  depends_on "bash-completion@2" => :optional

  def install
    bin.install "edkd"
  end
end
