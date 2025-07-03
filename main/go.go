components {
  id: "main"
  component: "/main/main.script"
}
embedded_components {
  id: "spinemodel"
  type: "spinemodel"
  data: "spine_scene: \"/assets/spineboy/spineboy.spinescene\"\n"
  "default_animation: \"idle\"\n"
  "skin: \"\"\n"
  "material: \"/defold-spine/assets/spine.material\"\n"
  "create_go_bones: true\n"
  ""
  scale {
    x: 0.5
    y: 0.5
  }
}
