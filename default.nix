{ lib, buildGoModule }:
buildGoModule {
  pname = "zzz";
  version = "unstable";

  src = ./.;

  vendorHash = "sha256-qgda2JLT+rtf9b+3JKhvu0uuXGKssLC0YrMghn1QQ8I=";

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
