import { NextApiRequest } from "next"
import { Handler, RequestMethod } from './handler'

type Router = (req: NextApiRequest) => Handler

export default function router(
  routeTable: Map<RequestMethod, Handler>
): Router {
  return (req: NextApiRequest) => {
    const method: RequestMethod = (req.method as RequestMethod)

    const handler = routeTable.get(method);

    if (handler === undefined) throw new HandlerNotProvidedError(req);

    return handler;
  }
}

class HandlerNotProvidedError extends Error {
  constructor(req: NextApiRequest) {
    super(
      `Did not supply handler for ${req.url} and Request 
      Method ${req.method}`
    );
  }
}