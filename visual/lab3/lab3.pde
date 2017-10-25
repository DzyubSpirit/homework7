import processing.opengl.*;
 
float A = -2, B = 10;
 
float min_x=-10, max_x=10, step_x=1;
int count_x = ceil((max_x-min_x)/step_x)+1;
float min_z=-10, max_z=10, step_z=1;
int count_z = ceil((max_z-min_z)/step_z)+1;

PVector[] vecs = new PVector[count_x*count_z];
int dim = 250;
 
void setup() {
  size(500,500,OPENGL);
  for (int i=0; i<count_x; i++) {
    for (int j=0; j<count_z; j++) {
      float x = min_x+step_x*i;
      float z = min_z+step_z*j;
      float y = 1 - (pow(x-A, 2)/A+pow(z-B, 2)/B);
      float draw_x = (x-min_x)/(max_x-min_x)*dim;
      float draw_z = (z-min_z)/(max_z-min_z)*dim;
      float draw_y = y;
      vecs[i*count_z+j] = new PVector(draw_x,draw_y,draw_z);
    }
  }
}
 
void draw() {
  background(0);
  translate(width/2,height/2);
  scale(1,-1,1); // so Y is up, which makes more sense in plotting
  rotateY(radians(frameCount));
 
  noFill();
  strokeWeight(1);
  box(dim);
 
  translate(-dim/2,-dim/2,-dim/2);
  for (int i=0; i<vecs.length; i++) {
    PVector v = vecs[i];
    stroke(255,75);
    strokeWeight(2);
    line(v.x,0,v.z,v.x,v.y,v.z);
    stroke(255);
    strokeWeight(5);
    point(v.x,v.y,v.z);
  }
}