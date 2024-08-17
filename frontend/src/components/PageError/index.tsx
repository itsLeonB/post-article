import { PageErrorProps } from './type';

const PageError: React.FC<PageErrorProps> = ({ error }) => {
  return <h1>{error}</h1>;
};

export default PageError;
