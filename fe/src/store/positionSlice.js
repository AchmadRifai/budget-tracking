import { createSlice } from '@reduxjs/toolkit'

export const positionSlice = createSlice({
    name: 'position',
    initialState: {
        position: 'login',
        auth: localStorage.getItem('auth') || '',
        role: ''
    },
    reducers: {
        setPosition(state, action) {
            state.position = action.payload
        },
        setAuth(state, action) {
            state.auth = action.payload
        },
        setRole(state, action) {
            state.role = action.payload
        }
    }
})

export const { setAuth, setPosition, setRole } = positionSlice.actions

export default positionSlice.reducer