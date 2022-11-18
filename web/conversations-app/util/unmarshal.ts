export default async function unmarshall<T>(res: Response): Promise<T> {
  const serializedRes = await res.json();
  return JSON.parse(serializedRes) as T;
}