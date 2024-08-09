import axios from 'axios'

export const delBudget = async (auth, id) => {
    const res = await axios.delete('/api/dashboard/budget/' + id, { headers: { 'Authorization': `Basic ${auth}`, "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    return res.data
}

export const editBudget = async (auth, name, amount, id) => {
    const res = await axios.put('/api/dashboard/budget/' + id, { name, amount }, { headers: { 'Authorization': `Basic ${auth}`, "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    return res.data
}

export const addBudget = async (auth, name, amount) => {
    const res = await axios.post('/api/dashboard/budget', { name, amount }, { headers: { 'Authorization': `Basic ${auth}`, "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    return res.data
}

export const budget = async (auth) => {
    const res = await axios.get('/api/dashboard/budget', { headers: { 'Authorization': `Basic ${auth}`, "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    return res.data
}

export const getChart = async (auth) => {
    const res = await axios.get('/api/dashboard/chart', { headers: { 'Authorization': `Basic ${auth}`, "Content-Type": "application/json" } })
    if (res.status !== 200) throw new Error(`${res.status} is ${res.statusText}`)
    return res.data
}

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