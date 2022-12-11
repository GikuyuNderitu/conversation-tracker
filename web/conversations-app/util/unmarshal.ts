export abstract class Prototype<T> {
  abstract construct(t: T): T;
}

export interface Shaper<T> {
  shape(t: T): T;
}

export default async function unmarshall<T>(res: Response, shaper?: Shaper<T>): Promise<T> {
  const serializedRes = await res.json();
  const parsed = JSON.parse(serializedRes) as T;

  if (shaper == undefined) return parsed;

  return shaper.shape(parsed);
}