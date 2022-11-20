import { NextApiResponse, NextApiRequest } from 'next'
import router from '../../util/api/router'

const _router = router(new Map().set('POST', createConvo));

export default function handler(
  req: NextApiRequest, res: NextApiResponse,
) {
  const _handler = _router(req);
  _handler(req, res);
}

function createConvo(req: NextApiRequest, res: NextApiResponse) {
  const title = req.body.title;

  // TODO(GikuyuNderitu): Plumb this request to NotesFe

  res.status(201).json({ title })
}
