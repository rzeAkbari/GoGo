function Login() {
  return (
    <div className='w-full'>
      <h1 className='text-center my-4'>{'Sign In'}</h1>
      <form className=''>
        <input
          className='placeholder:italic placeholder:text-slate-400 block bg-white w-full border
                   border-slate-300 rounded-md py-2 pl-9 pr-3 shadow-sm focus:outline-none my-10
                    focus:border-secondary focus:ring-secondary focus:ring-1 sm:text-sm'
          placeholder='user name ...'
          type='text'
          name='userName'
          aria-label='userName'
        />
        <input
          className='placeholder:italic placeholder:text-slate-400 block bg-white w-full border
                     border-slate-300 rounded-md py-2 pl-9 pr-3 shadow-sm focus:outline-none my-10
                     focus:border-secondary focus:ring-secondary focus:ring-1 sm:text-sm'
          placeholder='password ...'
          type='password'
          name='password'
          aria-label='password'
        />
        <button
          type='submit'
          className='block shadow-sm bourder border-slate-300 rounded-md p-4 my-10 bg-secondary
                     hover:bg-primary w-full'
        >
          Login
        </button>
      </form>
    </div>
  );
}

export default Login;
