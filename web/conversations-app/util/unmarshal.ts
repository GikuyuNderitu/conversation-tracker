export abstract class Prototype<T> {
  abstract construct(t: T): T;
}

export interface Shaper<T> {
  shape(t: T): T;
}

export default async function unmarshall<T>(res: Response, shaper?: Shaper<T>): Promise<T> {
  const json = await res.json();

  console.log(json)

  if (shaper == undefined) return json;

  return shaper.shape(json);
}