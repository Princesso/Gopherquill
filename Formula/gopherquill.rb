class Gopherquill < Formula
    desc "A powerful code review tool for developers"
    homepage "https://github.com/Princesso/gopherquill"
    url "https://github.com/Princesso/gopherquill/archive/refs/tags/v0.1.1.tar.gz"
    sha256 "468906a4bc3b65f9ab7637625d3f108c723f04b6"
  
    depends_on "go" => :build
  
    def install
      system "go", "build", "-o", bin/"gopherquill"
    end
  
    test do
      system "#{bin}/gopherquill", "--help"
    end
  end
  