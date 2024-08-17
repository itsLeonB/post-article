import { Provider } from 'react-redux';
import { RouterProvider } from 'react-router-dom';
import { createRouter } from './routes';
import { store } from '../store';

const AppRouter = () => {
  const router = createRouter();

  return <RouterProvider router={router} />;
};

function App() {
  return (
    <Provider store={store}>
      <AppRouter />
    </Provider>
  );
}

export default App;
