class Gopherquill < Formula
    desc "A powerful code review tool for developers"
    homepage "https://github.com/Princesso/gopherquill"
    url "https://github.com/Princesso/gopherquill/archive/refs/tags/v0.1.2.tar.gz"
    sha256 "7cb5cd0b546b59f3fd9b8369e28e7e874d09a38d30b7b5a9a1b6c18164cea15a"
  
    depends_on "go" => :build
  
    def install
      system "go", "build", "-o", bin/"gopherquill"
    end
  
    test do
      system "#{bin}/gopherquill", "--help"
    end

  end
  