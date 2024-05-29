import { useEffect } from 'react';
import {BrowserRouter, Route, Routes} from 'react-router-dom'
import { Auth } from './components/Auth'
import { User } from './components/User'
import axios from 'axios'
import { CsrfToken } from './types'

function App() {
  // 初回レンダリング時にcsrf取得
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token // ヘッダーに付与
    }
    getCsrfToken()
  }, [])
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Auth />} />
        <Route path="/user" element={<User />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
