const arr = [1, 2, 3];

arr.forEach((item, i) => {
  // arr.push(item * 2);
  arr.splice(i, 1);
  console.log(item);
});

const arr1 = [, , 3];

arr1.forEach((item) => {
  console.log(item);
});
