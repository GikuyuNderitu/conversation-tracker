import { NextApiResponse, NextApiRequest } from 'next'
import { json } from 'stream/consumers';
import router from '../../util/api/router'

const _router = router(new Map().set('POST', createConvo));

export default function handler(
  req: NextApiRequest, res: NextApiResponse,
) {
  const _handler = _router(req);
  _handler(req, res);
}

async function createConvo(req: NextApiRequest, res: NextApiResponse) {
  const title = req.body.title;

  const result = await fetch(
    'http://localhost:1337/convos', {
    method: 'POST',
    body: JSON.stringify({ title }),
  }
  );

  const json = await result.json();

  res.status(201).json(json)
}
