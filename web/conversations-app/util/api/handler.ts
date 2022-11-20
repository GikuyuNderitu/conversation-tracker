import { NextApiRequest, NextApiResponse } from "next";

export type RequestMethod = 'GET' | 'POST' | 'PUT' | 'DELETE';

export type Handler = (req: NextApiRequest, res: NextApiResponse) => void