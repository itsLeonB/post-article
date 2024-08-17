import { createBrowserRouter } from 'react-router-dom';
import Listing from './Listing';

export const createRouter = () =>
  createBrowserRouter([
    {
      path: '/',
      Component: () => <Listing />
    }
  ]);
