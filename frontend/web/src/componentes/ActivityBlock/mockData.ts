 export const activityData =[
    [{id:1,data:"1 Jan", contributions:8},{id:2,data:"2 Jan", contributions:0},{id:3,data:"3 Jan", contributions:3},{id:4,data:"4 Jan", contributions:1},{id:5,data:"5 Jan", contributions:6},{id:6,data:"6 Jan", contributions:3},{id:7,data:"7 Jan", contributions:0},],
 ]


let subArray = activityData[0];

for (var i = 0; i < 50; i++) {
  var newSubArray = subArray.map(function(obj) {
    return {...obj, id: obj.id + i * subArray.length, contributions:Math.floor(Math.random()*6)};
  });
  activityData.push(newSubArray);
}