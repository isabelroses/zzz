{ lib, buildGoModule }:
buildGoModule {
  pname = "zzz";
  version = "unstable";

  src = ./.;

  vendorHash = "sha256-PUxVEfmkUYZNKFefuWM5gaU8pdAPK8ZP1b5hx+Zj2GU=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "Code snippets in your terminal 🛌";
    mainProgram = "zzz";
    homepage = "https://github.com/isabelroses/zzz";
    license = lib.licenses.mit;
    maintainers = with lib.maintainers; [ isabelroses ];
  };
}
