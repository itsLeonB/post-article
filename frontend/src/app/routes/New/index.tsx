import { useNavigate } from 'react-router-dom';
import { Suspense } from 'react';
import { useAppDispatch } from '../../../hooks/useAppDispatch';
import Loader from '../../../components/Loader';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Inputs } from './type';
import { Post } from '../../../types/post';
import { toast, ToastContainer } from 'react-toastify';

import style from './style.module.css';
import { createPost } from '../../../store/thunks/postCreateThunk';

export const New: React.FC = () => {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<Inputs>();

  const publish: SubmitHandler<Inputs> = async (formData) => {
    const creatingPost: Post = {
      id: 0,
      title: formData.title,
      content: formData.content,
      category: formData.category,
      status_id: 1
    };

    const res = await dispatch(createPost(creatingPost)).unwrap();

    if (res.error) {
      toast.error(res.error.message, { theme: 'dark' });
    } else {
      navigate('/');
    }
  };

  const draft: SubmitHandler<Inputs> = async (formData) => {
    const creatingPost: Post = {
      id: 0,
      title: formData.title,
      content: formData.content,
      category: formData.category,
      status_id: 2
    };

    const res = await dispatch(createPost(creatingPost)).unwrap();

    if (res.error) {
      toast.error(res.error.message, { theme: 'dark' });
    } else {
      navigate('/');
    }
  };

  return (
    <Suspense fallback={<Loader />}>
      <ToastContainer />
      <form className={style.form}>
        <label className={style.label}>
          Title:
          <input
            type="text"
            {...register('title', {
              required: true,
              minLength: 20,
              maxLength: 200
            })}
          />
          {<span>{errors.title?.type}</span>}
        </label>
        <label className={style.label}>
          Content:
          <textarea
            rows={5}
            {...register('content', { required: true, minLength: 200 })}
          />
          {<span>{errors.content?.type}</span>}
        </label>
        <label className={style.label}>
          Category:
          <input
            type="text"
            {...register('category', {
              required: true,
              minLength: 3,
              maxLength: 100
            })}
          />
          {<span>{errors.category?.type}</span>}
        </label>
        <button onClick={handleSubmit(publish)}>Publish</button>
        <button onClick={handleSubmit(draft)}>Draft</button>
      </form>
    </Suspense>
  );
};
