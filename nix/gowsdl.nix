{
  fetchFromGitHub,
  installShellFiles,
  makeWrapper,
  makeBinPath,
  go,
  goBuild,
  ...
}:
goBuild rec {
  pname = "gowsdl";
  version = "0.3.1";

  src = fetchFromGitHub {
    owner = "hooklift";
    repo = "gowsdl";
    rev = "v${version}";
    hash = "sha256-EnEi8FECUUZcLyfRGj6Xion+WxSRPrjhryK4xjaJlWI=";
  };

  vendorHash = null;

  subPackages = [ "cmd/gowsdl" ];

  nativeBuildInputs = [
    installShellFiles
    makeWrapper
  ];

  ldflags = [
    "-s"
    "-X main.version=${version}"
    "-X main.commit=v${version}"
    "-X main.date=19700101-00:00:00"
  ];

  # postInstall = ''
  #   for shell in bash zsh fish; do
  #     HOME=$TMPDIR $out/bin/gowsdl completion $shell > gowsdl.$shell
  #     installShellCompletion gowsdl.$shell
  #   done
  # '';

  allowGoReference = true;

  postFixup = ''
    wrapProgram $out/bin/gowsdl --prefix PATH : ${makeBinPath [ go ]}
  '';
}
