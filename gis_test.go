package gis

import (
  "testing"
  "log"
  "github.com/omniscale/go-proj"
)

func TestCheckGisDistance(t *testing.T) {
  d1 := Distance(55.6553, 37.7581, 55.6572, 37.759)
  if d1 != 218.9277172080462 {
    t.Error(
      "For", "Distance 1",
      "expected", 218.9277172080462,
      "got", d1,
    )
  }

  d2 := Distance(55.657, 37.7605, 55.6573, 37.7604)
  if d2 != 33.98099254965385 {
    t.Error(
      "For", "Distance 2",
      "expected", 33.98099254965385,
      "got", d2,
    )
  }

  d3 := Distance(55.6575, 37.7601, 55.6551, 37.7577)
  if d3 != 306.7484932040614 {
    t.Error(
      "For", "Distance 3",
      "expected", 306.7484932040614,
      "got", d3,
    )
  }

}

/*
 def coord_translate(light_id):
  lights_dict = {
            '8509CD6A': [55.656066, 37.758993],
            '41737B5E': [55.656964, 37.754895],
            'C7B5997B': [55.656416, 37.757949],
            '9C684B85': [55.656269, 37.757798],
            '41239DE5': [55.656267, 37.759387],
            '59AC23C5': [55.656664, 37.756162],
            'DD48803D': [55.656646, 37.755691],
            '26A40093': [55.656543, 37.75827],
            '650AB959': [55.655809, 37.75518],
            'E8F2B455': [55.656038, 37.758997],
            '7D91EFDA': [55.657801, 37.758934],
            'C63258FA': [55.656468, 37.754454],
            'DA3C6449': [55.656637, 37.756216],
            'F1DC7775': [55.657292, 37.756131],
            'D1DA1B1D': [55.65607, 37.75799],
            'BC8A0AC5': [55.65591, 37.758305],
            '94937545': [55.656222, 37.759331],
            '29832648': [55.657008, 37.756464],
            '1E3E188F': [55.657105, 37.756931],
            'EA1AC040': [55.656901, 37.755333],
            '964FB3AC': [55.656936, 37.754872],
            'B1E62E2A': [55.656818, 37.754598],
            'AECFB55E': [55.656285, 37.757827],
            '649FFEA4': [55.657306, 37.755784],
            '006F3293': [55.656918, 37.756557],
            '345427EF': [55.656307, 37.754651],
            'C8FF6C12': [55.65625, 37.758824],
            '6AF79A65': [55.657382, 37.759049],
            '5837818E': [55.657449, 37.758783],
            '05A03684': [55.656461, 37.758518],
            'F6523526': [55.657142, 37.756633],
            '0D3F0D59': [55.65579, 37.75858],
            'AA6F3CAD': [55.656483, 37.758169],
            '139FC823': [55.657224, 37.756536],
            'D36DA6D4': [55.656088, 37.758063],
            '01EF5783': [55.656931, 37.755247],
            'CBC1476E': [55.656591, 37.758159],
            'DAB77FAB': [55.656883, 37.756819],
            '866117D4': [55.656207, 37.758813],
            'B8B0F483': [55.657163, 37.755153],
            'FCC408CB': [55.655821, 37.75533],
            '06E188FB': [55.655664, 37.756842],
            'A94AA382': [55.657284, 37.756347],
            '82F57D7A': [55.655899, 37.758244],
            '19EF2100': [55.655855, 37.755184],
            'C2C361AC': [55.656204, 37.75566],
            'E7E98AC6': [55.657671, 37.759005],
            '19135278': [55.656328, 37.758682],
            '00756B27': [55.656435, 37.75446],
            '7078305B': [55.65593, 37.75887],
            '64F39D2F': [55.65673, 37.756537],
            'E2464BBF': [55.656893, 37.756871],
            'F0B038E9': [55.655377, 37.757953],
            'FBC8382A': [55.656837, 37.758816],
            'CF1B01A3': [55.657043, 37.758318],
            '879CD6AA': [55.657151, 37.755556],
            '011CCC34': [55.655976, 37.756294],
            'B611BB41': [55.655702, 37.758105],
            'AA29350E': [55.658348, 37.758311],
            'FEFB0613': [55.656329, 37.755461],
            '28DD4F36': [55.656171, 37.760225],
            'F15DBF98': [55.655158, 37.756635],
            '127FCA46': [55.655478, 37.758576],
            'C88ECA81': [55.655822, 37.757952],
            '81339313': [55.657771, 37.758488],
            '23263AED': [55.656268, 37.755296],
            'E4029871': [55.656649, 37.760168],
            'BB4E2D9F': [55.656934, 37.755494],
            'BB1E5D8E': [55.655783, 37.756072],
            '7137043F': [55.655905, 37.757758],
            'B23F544C': [55.656443, 37.755694],
            '57EEEAAE': [55.656402, 37.757289],
            '69214276': [55.657474, 37.759197],
            'E61B1A21': [55.656671, 37.754525],
            'DAC4755C': [55.65685, 37.754571],
            '9B533977': [55.656521, 37.759818],
            'CAB93978': [55.656801, 37.756791],
            '5D1D32A9': [55.656921, 37.755534],
            'C071A622': [55.656673, 37.755775],
            '63254EF2': [55.656457, 37.758729],
            '3967D0C8': [55.656895, 37.759845],
            'DDEB25AF': [55.657232, 37.759467],
            'B74698F5': [55.656222, 37.75733],
            '98792FE7': [55.655818, 37.755865],
            'B20A9C2A': [55.654999, 37.756129],
            '93F82376': [55.657586, 37.757042],
            '468BC761': [55.657093, 37.759679],
            'C529EB9F': [55.656642, 37.754559],
            '2E93F840': [55.656789, 37.756726],
            'FE58646B': [55.657313, 37.756002],
            '3AF911FC': [55.656023, 37.755539],
            '7190EF24': [55.656922, 37.757344],
            'A2C82CBA': [55.65682, 37.757706],
            '9B35EB7B': [55.655609, 37.758765],
            'D9A01D76': [55.656037, 37.756569],
            '0275DC36': [55.657238, 37.75501],
            '4D8ACB65': [55.655331, 37.756592],
            '8AAD83F3': [55.656478, 37.755453],
            '2AF6AC46': [55.656436, 37.755738],
            'F7A83A24': [55.657661, 37.757764],
            '7865B1B8': [55.65615, 37.754805],
            '0821042F': [55.65638, 37.757195],
            'A83A2AC8': [55.65676, 37.756501],
            'B3FA9B20': [55.655744, 37.757226],
            '6278FFA7': [55.658104, 37.757824],
            '5EC93D16': [55.656256, 37.755969],
            '59397C15': [55.657454, 37.757972],
            '7E4CEA1A': [55.657297, 37.758003],
            'B9D69503': [55.655502, 37.756393],
            '8C6AF889': [55.65531, 37.758098],
            'F55DF71F': [55.656476, 37.755511],
            '8D79622C': [55.65565, 37.7574],
            'B220C0F3': [55.655126, 37.757544],
            'A7012F64': [55.657274, 37.755174],
            '53744C8C': [55.655939, 37.759384],
            '8F2A9145': [55.657309, 37.755091],
            'BBC8A534': [55.656462, 37.756078],
            '945142B9': [55.657301, 37.756911],
            'BEDEC3D2': [55.656227, 37.75686],
            '9049A5C4': [55.65596, 37.759807],
            'BADDE4E6': [55.655057, 37.757112],
            'F0A32120': [55.65769, 37.758751],
            '6F059F40': [55.65464, 37.756531],
            '062E7C8D': [55.65712, 37.75715],
            '4CF71F03': [55.656167, 37.760171],
            '8189CDD9': [55.656458, 37.756092],
            'B654BC64': [55.656584, 37.759198],
            '8F784BE0': [55.654737, 37.756354],
            '5A0F535E': [55.656137, 37.756712],
            '641D7C92': [55.658478, 37.758654],
            '401F79A8': [55.656139, 37.756667],
            '5A4C740C': [55.657492, 37.757692],
            '52FDF8D8': [55.657769, 37.758239],
            'BA64274B': [55.656118, 37.755524],
            '96D93D9D': [55.65574, 37.7592],
            'FE47993D': [55.655737, 37.757552],
            '38AC27AF': [55.6579, 37.757459],
            '7CE48ED6': [55.65629, 37.757024],
            '61002605': [55.656596, 37.756385],
            'D46C42E9': [55.656641, 37.756376],
            '6EC268B2': [55.656383, 37.75948],
            'AA25EDB5': [55.65624, 37.75523],
            '414E4C4F': [55.658075, 37.758633],
            '6A7C45C6': [55.656705, 37.758287],
            '51B05E62': [55.655355, 37.758127],
            '34EC58C5': [55.654982, 37.756792],
            '011CCC34': [55.655977, 37.756293] }
 */
 
func TestCheckGisEPSG2WGS84(t *testing.T) {
  //var lonlat = []float64{16243.080849, -1200.534339} // 55.655977, 37.756293

  wgs84, err := proj.NewEPSG(4326) // WGS 84 (EPSG:4326) 
  if err != nil {
    log.Printf("ERR: COORD: %v\n", err)
  }

  // proj by proj4 definition string
  //msk, err := proj.New("+proj=tmerc +lat_0=55.6666666667 +lon_0=37.5 +x_0=0 +y_0=0 +k_0=1. +a=6377397 +rf=299.15 +towgs84=396,165,557.7,-0.05,0.04,0.01,0 +no_defs")
  msk, err := proj.New("+proj=merc +a=6378137 +b=6378137 +lat_ts=0.0 +lon_0=0.0 +x_0=0.0 +y_0=0 +k=1.0 +units=m +nadgrids=@null +wktext  +no_defs")
  if err != nil {
    log.Printf("ERR: COORD: %v\n", err)
  }
  var xs = []float64{16243.080849} // 55.655977, 37.756293
  var ys = []float64{-1200.534339} // 55.655977, 37.756293
                                   // 55.66,     37.76
  // transform all coordinates to WGS84 (in-place)
  if err := msk.Transform(wgs84, xs, ys); err != nil {
    log.Printf("ERR: COORD: %v\n", err)
  }
 	log.Printf("COORD: %.2f, %.2f\n", xs[0], ys[0])
  
  
  // 16154.6809, -1090.970398 => 55.656964, 37.754895
}
