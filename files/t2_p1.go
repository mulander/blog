package main
import (
  "os";
  "fmt";
  "math";
  "image";
  "image/png"
)

const ResX = 1024;
const ResY = 768;
const MaxPrimCount = 64;
const MaxLightCount = 10;
const MaxThreads = 4;

type Vector3D struct {
  x, y, z, w float64;
}
type Material struct {
  Specular, Diffusive, Reflective float64;
  Color Vector3D;
}

type PrimSphere struct {
  Position Vector3D;
  Radius float64;
  m *Material;
}

type Ray struct {
  Direction, Origin Vector3D;
}

type Light struct {
  Position, Color Vector3D;
}





type GlobalsStr struct {
  img *image.RGBA;

  PrimitiveCount int;
  PrimitiveList [MaxPrimCount]*PrimSphere;

  LightCount int;
  LightList [MaxLightCount]*Light;

}

var Globals GlobalsStr


func main() {
  fmt.Printf("Simple RT by gynvael.coldwind//vx (http://gynvael.coldwind.pl)\n");

  fmt.Printf("Creating scene...\n");

  Globals.img = image.NewRGBA(ResX, ResY);

  var Mirror, Green, Red Material;
  { Mirror.Color.x = 0.6; Mirror.Color.y = 0.6; Mirror.Color.z = 0.6; };
  Mirror.Specular = 0.3;
  Mirror.Diffusive = 0.2;
  Mirror.Reflective = 0.8;

  { Green.Color.x = 0.1; Green.Color.y = 1.0; Green.Color.z = 0.1; };
  Green.Specular = 0.1;
  Green.Diffusive = 0.3;
  Green.Reflective = 0.4;

  { Red.Color.x = 1.0; Red.Color.y = 0.1; Red.Color.z = 0.1; };
  Red.Specular = 0.1;
  Red.Diffusive = 0.3;
  Red.Reflective = 0.4;

  var SpherePosMap string;

  SpherePosMap = 
    "........." +
    ".ggg....." +
    ".g...rrr." +
    ".g.g.r.r." +
    ".ggg.rrr." +
    ".........";

  SpherePosMap = SpherePosMap;
  p := 0;


  var SpherePos Vector3D;
  for j := 0; j < 6; j++ {
    for i := 0; i < 9; i++ {
      m := &Mirror;

      z := float64(2.0);
      sn := math.Sin(float64(i+j)) * 0.8;

      switch
      {
        case SpherePosMap[p] == 'g': z += -0.5; m = &Green;
        case SpherePosMap[p] == 'r': z += -0.5; m = &Red;
        default: z += sn;
      }

      { SpherePos.x = -2.0 + float64(i) * 0.5; SpherePos.y = 1.25 - float64(j) * 0.5; SpherePos.z = z; };
      AddSphere(&SpherePos, 0.25, m);
      p++;
    }
  }

  var LightPos Vector3D; LightPos.x = 0.0; LightPos.y = 0.0; LightPos.z = 0.0;;
  var LightColor Vector3D; LightColor.x = 2.0; LightColor.y = 2.0; LightColor.z = 2.0;;
  AddLight(&LightPos, &LightColor);

  fmt.Printf("Rendering...\n");


  c := make(chan int);
  for i := 0; i < MaxThreads; i++ {
    fmt.Printf("[%d] Thread start\n", i);
    go Render(Globals.img, MaxThreads, i, c);
  }


  for i := 0; i < MaxThreads; i++ {
    fmt.Printf("[%d] Thread finished\n", <-c);
  }


  fmt.Printf("Writing test.png image...\n");
  fd, err:= os.Open("test.png", os.O_WRONLY | os.O_CREATE, 0666);

  if(err != nil) {
    fmt.Printf("error: %s\n", err.String());
    return
  }

  png.Encode(fd, Globals.img);

  fmt.Printf("Done.\n");
}

func AddLight(Pos, Color *Vector3D) {
  if Globals.LightCount < MaxLightCount {
    l := new(Light);
    { l.Position.x = Pos.x; l.Position.y = Pos.y; l.Position.z = Pos.z; };
    { l.Color.x = Color.x; l.Color.y = Color.y; l.Color.z = Color.z; };
    Globals.LightList[Globals.LightCount] = l;
    Globals.LightCount++;
  }
}


func AddSphere(Pos *Vector3D, Rad float64, m *Material) {
  if Globals.PrimitiveCount < MaxPrimCount {
    p := new(PrimSphere);
    { p.Position.x = Pos.x; p.Position.y = Pos.y; p.Position.z = Pos.z; };
    p.Radius = Rad;
    p.m = m;
    Globals.PrimitiveList[Globals.PrimitiveCount] = p;
    Globals.PrimitiveCount++;
  }
}

func (this *PrimSphere) Intersect(ray *Ray, dist *float64) int {
  var v_precalc Vector3D;
  var det_precalc, b, det, i1, i2 float64;

  { v_precalc.x = ray.Origin.x; v_precalc.y = ray.Origin.y; v_precalc.z = ray.Origin.z; };
  { v_precalc.x -= this.Position.x; v_precalc.y -= this.Position.y; v_precalc.z -= this.Position.z; };
  det_precalc = this.Radius * this.Radius - (v_precalc.x*v_precalc.x + v_precalc.y*v_precalc.y + v_precalc.z*v_precalc.z);

  b = - (v_precalc.x*ray.Direction.x + v_precalc.y*ray.Direction.y + v_precalc.z*ray.Direction.z);

  det = b*b + det_precalc;

  retval := int(0);

  if(det > 0) {
    det = math.Sqrt(det);
    i1 = b - det;
    i2 = b + det;
    switch {
      case i2 > 0 && i1 < 0:
        retval = -1;
        *dist = i2;
      case i2 > 0 && i1 >= 0:
        retval = 1;
        *dist = i1;
    }
  }

  return retval;
}

func (this *PrimSphere) Normal(Ret, pos *Vector3D) {
  { Ret.x = pos.x; Ret.y = pos.y; Ret.z = pos.z; };
  { Ret.x -= this.Position.x; Ret.y -= this.Position.y; Ret.z -= this.Position.z; };
  f := float64(1.0 / this.Radius);
  { Ret.x *= f; Ret.y *= f; Ret.z *= f; };
  { _l := float64(1.0 / (math.Sqrt(Ret.x*Ret.x+Ret.y*Ret.y+Ret.z*Ret.z))); Ret.x *= _l; Ret.y *= _l; Ret.z *= _l; };
}


func Trace(RetVector *Vector3D, ray *Ray, refl_depth int) {
  var color Vector3D; color.x = 0.02; color.y = 0.1; color.z = 0.17;;


  dist := float64(1000000000.0);
  var prim *PrimSphere = nil;

  for i := int(0); i < Globals.PrimitiveCount; i++ {
    var temp_dist float64;

    p := Globals.PrimitiveList[i];

    res := p.Intersect(ray, &temp_dist);

    if res == 0 {
      continue;
    }

    if temp_dist < dist {
      prim = p;
      dist = temp_dist;

    }
  }

  dist = dist;
  prim = prim;

  if prim == nil {
    { RetVector.x = color.x; RetVector.y = color.y; RetVector.z = color.z; };
    return;
  }


  var pi Vector3D;
  { pi.x = ray.Direction.x; pi.y = ray.Direction.y; pi.z = ray.Direction.z; };
  { pi.x *= dist; pi.y *= dist; pi.z *= dist; };
  { pi.x += ray.Origin.x; pi.y += ray.Origin.y; pi.z += ray.Origin.z; };

  var prim_color Vector3D;
  { prim_color.x = prim.m.Color.x; prim_color.y = prim.m.Color.y; prim_color.z = prim.m.Color.z; };


  for i := 0; i < Globals.LightCount; i++ {
    lightiter := Globals.LightList[i];

    var L, N Vector3D;
    { L.x = lightiter.Position.x; L.y = lightiter.Position.y; L.z = lightiter.Position.z; };
    { L.x -= pi.x; L.y -= pi.y; L.z -= pi.z; };

    { _l := float64(1.0 / (math.Sqrt(L.x*L.x+L.y*L.y+L.z*L.z))); L.x *= _l; L.y *= _l; L.z *= _l; };

    prim.Normal(&N, &pi);

    if prim.m.Diffusive > 0.0 {
      dot := (L.x*N.x + L.y*N.y + L.z*N.z);
      if dot > 0.0 {
        diff := dot * prim.m.Diffusive;

        var color_add Vector3D;
        { color_add.x = lightiter.Color.x; color_add.y = lightiter.Color.y; color_add.z = lightiter.Color.z; };
        { color_add.x *= prim_color.x; color_add.y *= prim_color.y; color_add.z *= prim_color.z; };
        { color_add.x *= diff; color_add.y *= diff; color_add.z *= diff; };
        { color.x += color_add.x; color.y += color_add.y; color.z += color_add.z; };


      }
    }

    if prim.m.Specular > 0.0 {



      var R1, R Vector3D;

      { R1.x = N.x; R1.y = N.y; R1.z = N.z; };
      R2 := (L.x*N.x + L.y*N.y + L.z*N.z) * 2.0;
      { R1.x *= R2; R1.y *= R2; R1.z *= R2; };
      { R.x = L.x; R.y = L.y; R.z = L.z; };
      { R.x -= R1.x; R.y -= R1.y; R.z -= R1.z; };

      dot := (ray.Direction.x*R.x + ray.Direction.y*R.y + ray.Direction.z*R.z);
      if dot > 0 {

        dot *= dot; dot *= dot; dot *= dot;
        spec := dot * prim.m.Specular;

        var color_add Vector3D;
        { color_add.x = lightiter.Color.x; color_add.y = lightiter.Color.y; color_add.z = lightiter.Color.z; };
        { color_add.x *= spec; color_add.y *= spec; color_add.z *= spec; };
        { color.x += color_add.x; color.y += color_add.y; color.z += color_add.z; };
      }
    }

  refl := prim.m.Reflective;
  if refl > 0.0 && refl_depth < 4 {

    var N, R, R1, newPi Vector3D;
    prim.Normal(&N, &pi);


    { R.x = ray.Direction.x; R.y = ray.Direction.y; R.z = ray.Direction.z; };
    { R1.x = N.x; R1.y = N.y; R1.z = N.z; };

    R2 := (ray.Direction.x*N.x + ray.Direction.y*N.y + ray.Direction.z*N.z) * 2.0;
    { R1.x *= R2; R1.y *= R2; R1.z *= R2; };
    { R.x -= R1.x; R.y -= R1.y; R.z -= R1.z; };


    var rcol Vector3D; rcol.x = 0.0; rcol.y = 0.0; rcol.z = 0.0;;



    { newPi.x = R.x; newPi.y = R.y; newPi.z = R.z; };
    { newPi.x *= 0.0001; newPi.y *= 0.0001; newPi.z *= 0.0001; };
    { newPi.x += pi.x; newPi.y += pi.y; newPi.z += pi.z; };

    var tempr Ray;
    { tempr.Origin.x = newPi.x; tempr.Origin.y = newPi.y; tempr.Origin.z = newPi.z; };
    { tempr.Direction.x = R.x; tempr.Direction.y = R.y; tempr.Direction.z = R.z; };


    Trace(&rcol, &tempr, refl_depth+1);


    { rcol.x *= refl; rcol.y *= refl; rcol.z *= refl; };
    { rcol.x *= prim_color.x; rcol.y *= prim_color.y; rcol.z *= prim_color.z; };
    { color.x += rcol.x; color.y += rcol.y; color.z += rcol.z; };
  }


  }

  { RetVector.x = color.x; RetVector.y = color.y; RetVector.z = color.z; };
}

func Render(img *image.RGBA, ThreadCount, ThreadId int, c chan int) {
  var CameraPos Vector3D; CameraPos.x = 0.0; CameraPos.y = 0.0; CameraPos.z = -5.0;;
  var WX1, WX2, WY1, WY2 float64 = -2.0, 2.0, 1.5, -1.5;
  var b = img.Bounds();

  DX := float64((WX2 - WX1) / float64(b.Max.X));
  DY := float64((WY2 - WY1) / float64(b.Max.Y));

  SX := WX1; SY := WY1 + DY * float64(ThreadId);

  var x, y int;



  for y = ThreadId; y < b.Max.Y; y += ThreadCount {
    SX = WX1;

    for x = b.Min.X; x < b.Max.X; x++ {
      var CameraTarget Vector3D; CameraTarget.x = SX; CameraTarget.y = SY; CameraTarget.z = 0.0;;

      var ray Ray;
      { ray.Origin.x = CameraPos.x; ray.Origin.y = CameraPos.y; ray.Origin.z = CameraPos.z; };
      { ray.Direction.x = CameraTarget.x; ray.Direction.y = CameraTarget.y; ray.Direction.z = CameraTarget.z; };
      { ray.Direction.x -= ray.Origin.x; ray.Direction.y -= ray.Origin.y; ray.Direction.z -= ray.Origin.z; };
      { _l := float64(1.0 / (math.Sqrt(ray.Direction.x*ray.Direction.x+ray.Direction.y*ray.Direction.y+ray.Direction.z*ray.Direction.z))); ray.Direction.x *= _l; ray.Direction.y *= _l; ray.Direction.z *= _l; };

      var color Vector3D;
      Trace(&color, &ray, 0);


      var r, g, b int;
      r = int(color.x * 255.0);
      g = int(color.y * 255.0);
      b = int(color.z * 255.0);

      switch {
        case r > 255: r = 255;
        case r < 0: r = 0;
      }

      switch {
        case g > 255: g = 255;
        case g < 0: g = 0;
      }

      switch {
        case b > 255: b = 255;
        case b < 0: b = 0;
      }

      var cl image.RGBAColor;

      cl.R = uint8(r);
      cl.G = uint8(g);
      cl.B = uint8(b);
      cl.A = 255;

      img.Pix[y*img.Stride+x] = cl;



      SX += DX;
    }
    SY += DY * float64(ThreadCount);
  }


  c <-ThreadId;
}
