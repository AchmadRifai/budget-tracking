import axios from 'axios'

export const adminDeleteUser = async (auth, id) => {
    const res = await axios.delete('/api/dashboard/admin/users/' + id, { headers: { 'Authorization': `Basic ${auth}`, "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    return res.data
}

export const adminUsers = async (auth) => {
    const res = await axios.get('/api/dashboard/admin/users', { headers: { 'Authorization': `Basic ${auth}`, "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    return res.data
}

export const dashboardLogout = async (auth) => {
    const res = await axios.get('/api/dashboard/logout', { headers: { 'Authorization': `Basic ${auth}`, "Content-Type": "application/json" } })
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