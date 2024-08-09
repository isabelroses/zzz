{ lib, buildGoModule }:
buildGoModule {
  pname = "zzz";
  version = "unstable";

  src = ./.;

  vendorHash = "sha256-Iz5+jZs80wWlbQ6pw0/CQHy2gaFm9pT/LdipfW9Hg4o=";

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
