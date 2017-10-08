// The statements in the setup() function 
// execute once when the program begins
void setup() {
  size(640, 480);  // Size should be the first statement
}

void draw() {
  clear();
  background(255);
  float radius = 10 * (cos(angleFromTime()) * 0.5 + 1);  
  drawOrnamentCross(10, 10, 400, 30, radius);
}

float angleFromTime() {
  return (minute() + second() + millis()) / 500.0; 
}

void drawSign(float x, float y, float radius) {
  ellipse(x, y, radius*2, radius*2);
  float p = radius / sqrt(2);
  drawCross(x, y, p);
}

void drawCross(float x, float y, float size) {
  line(x-size, y-size, x+size, y+size);
  line(x-size, y+size, x+size, y-size);
}

void drawOrnamentCross(float lx, float ly, float size, float gap, float radius) {
  float p = ((size - 2 * radius) % gap) / 2;
  for (float x = p + radius; x < size; x += gap) {
    drawSign(lx + x, ly + x, radius);
    drawSign(lx + x, ly + size - x, radius);  
  }
}