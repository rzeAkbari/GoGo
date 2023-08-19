import { Fragment } from 'react';
import Login from './components/Login';
import logo from '/logo.png';

function App() {
  return (
    <Fragment>
      <header>
        <div className='w-full h-24 bg-primary'>
          <img src={logo} alt='logo' className='h-24' />
        </div>
      </header>
      <main>
        <section className='grid grid-cols-3'>
          <div className='col-start-2'>
            <Login />
          </div>
        </section>
      </main>
    </Fragment>
  );
}

export default App;
