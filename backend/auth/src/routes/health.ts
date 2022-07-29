import express from 'express';

const router = express.Router();

router.get('/api/users/health', (req, res) => {
  res.send({message: "Auth service is healthy!"});
});

export { router as healthRouter };
