import { createBrowserRouter } from 'react-router-dom';
import Listing from './Listing';

export const createRouter = () =>
  createBrowserRouter([
    {
      path: '/',
      element: <Listing />
    },
    {
      path: '/:id',
      lazy: async () => {
        const { Edit } = await import('./Edit');
        return { Component: () => <Edit /> };
      }
    }
  ]);
