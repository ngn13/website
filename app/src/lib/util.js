function click() {
  let audio = new Audio("/click.wav");
  audio.play();
}

let colors_pos = -1;
const colors = ["yellow", "cyan", "green", "pinkish", "red", "blue"];

function color() {
  if (colors_pos < 0) colors_pos = Math.floor(Math.random() * colors.length);
  else if (colors_pos >= colors.length) colors_pos = 0;

  return colors[colors_pos];
}

export { click, color };
