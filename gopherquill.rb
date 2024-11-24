class Gopherquill < Formula
    desc "A powerful code review tool for developers"
    homepage "https://github.com/Princesso/gopherquill"
    url "https://github.com/Princesso/gopherquill/archive/refs/tags/v0.1.2.tar.gz"
    sha256 "0a6c8831d314ce8bc079f3ce03e53674fcf11778"
  
    depends_on "go" => :build
  
    def install
      system "go", "build", "-o", bin/"gopherquill"
    end
  
    test do
      system "#{bin}/gopherquill", "--help"
    end
  end
  