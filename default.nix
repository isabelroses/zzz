{ lib, buildGoModule }:
buildGoModule {
  pname = "zzz";
  version = "unstable";

  src = ./.;

  vendorHash = "sha256-9Q2MwkZ6PBu+Kqg1YV4YFGJuXP4/Si+fOQAB57eKODU=";

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
