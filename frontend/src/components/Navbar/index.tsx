import { Link } from 'react-router-dom';
import style from './style.module.css';

const Navbar: React.FC = () => {
  return (
    <nav className={style.navbar}>
      <h2>Post Article</h2>
      <Link to="/">List</Link>
      <Link to="/new">New</Link>
      <Link to="/preview">Preview</Link>
    </nav>
  );
};

export default Navbar;
