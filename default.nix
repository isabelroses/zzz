{
  lib,
  buildGoModule,
}:
buildGoModule {
  pname = "zzz";
  version = "unstable";

  src = ./.;

  vendorHash = "sha256-ePHkrsc9NJO8c3J1eDkFeSLvVs5flSeiTmXqkfF261s=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "Code snippets in your terminal 🛌";
    mainProgram = "zzz";
    homepage = "https://github.com/isabelroses/zzz";
    license = lib.licenses.mit;
    maintainers = with lib.maintainers; [isabelroses];
  };
}
