import { useNavigate, useParams } from 'react-router-dom';
import { useAppSelector } from '../../../hooks/useAppSelector';
import { Suspense, useEffect } from 'react';
import { useAppDispatch } from '../../../hooks/useAppDispatch';
import { fetchPost } from '../../../store/thunks/postThunk';
import Loader from '../../../components/Loader';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Inputs } from './type';
import { Post } from '../../../types/post';
import { updatePost } from '../../../store/thunks/postUpdateThunk';
import { toast, ToastContainer } from 'react-toastify';
import { resetStatus } from '../../../store/slices/postSlice';

import style from './style.module.css';

export const Edit: React.FC = () => {
  const { id } = useParams();
  const dispatch = useAppDispatch();
  const { data, status, error } = useAppSelector((state) => state.post);
  const navigate = useNavigate();

  useEffect(() => {
    dispatch(resetStatus());
  }, [id, dispatch]);

  useEffect(() => {
    if (status === 'idle' && id) {
      dispatch(fetchPost(parseInt(id)));
    }
  }, [status, dispatch, id]);

  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors }
  } = useForm<Inputs>();

  const publish: SubmitHandler<Inputs> = async (formData) => {
    const updatingPost: Post = {
      id: parseInt(id!),
      title: formData.title,
      content: formData.content,
      category: formData.category,
      status_id: 1
    };

    await dispatch(updatePost(updatingPost));

    if (error) {
      toast.error(error, { theme: 'dark' });
    } else {
      navigate('/');
    }
  };

  const draft: SubmitHandler<Inputs> = async (formData) => {
    const updatingPost: Post = {
      id: parseInt(id!),
      title: formData.title,
      content: formData.content,
      category: formData.category,
      status_id: 2
    };

    await dispatch(updatePost(updatingPost));

    if (error) {
      toast.error(error, { theme: 'dark' });
    } else {
      navigate('/');
    }
  };

  useEffect(() => {
    if (data) {
      setValue('title', data.title);
      setValue('content', data.content);
      setValue('category', data.category);
    }
  }, [data]);

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
