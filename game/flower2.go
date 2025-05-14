components {
  id: "flower"
  component: "/game/flower.script"
  properties {
    id: "eat_animation"
    value: "Explosion"
    type: PROPERTY_TYPE_HASH
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"flower\"\n"
  "mask: \"player\"\n"
  "mask: \"bullet\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: -20.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 35.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "rivemodel"
  type: "rivemodel"
  data: "scene: \"/assets/flower1.rivescene\"\n"
  "default_animation: \"idle\"\n"
  "material: \"/defold-rive/assets/rivemodel.material\"\n"
  "artboard: \"Flower1\"\n"
  ""
}
