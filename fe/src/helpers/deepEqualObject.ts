export function deepEqualObject<T>(obj1: T, obj2: T): boolean {
  console.log(obj1, obj2);
  if (obj1 === obj2) {
    return true;
  }
  if (
    typeof obj1 !== 'object' ||
    typeof obj2 !== 'object' ||
    obj1 === null ||
    obj2 === null
  ) {
    return false;
  }
  const keys1 = Object.keys(obj1) as Array<keyof T>;
  const keys2 = Object.keys(obj2) as Array<keyof T>;
  if (keys1.length !== keys2.length) {
    return false;
  }
  for (const key of keys1) {
    if (!keys2.includes(key) || !deepEqualObject(obj1[key], obj2[key])) {
      return false;
    }
  }
  return true;
}
