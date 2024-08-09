import axios from 'axios'

export const dashboardLogout = async () => {
    const res = await axios.get('/api/dashboard/logout', { headers: { Authorization: `Basic ${localStorage.getItem('auth')}`, "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    return res.data
}

export const register = async (fullname, username, password, email) => {
    const res = await axios.post(`/api/register`, { username, password, 'full_name': fullname, email }, { headers: { "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    const header = { authorization: res.headers.getAuthorization() }
    const body = res.data
    return { header, body }
}

export const login = async (username, password) => {
    const res = await axios.post(`/api/login`, { username, password }, { headers: { "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    const header = { authorization: res.headers.getAuthorization() }
    const body = res.data
    return { header, body }
}