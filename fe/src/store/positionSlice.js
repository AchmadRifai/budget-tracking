import { createSlice } from '@reduxjs/toolkit'

export const positionSlice = createSlice({
    name: 'position',
    initialState: {
        position: 'login',
        auth: '',
        role: '',
        name: '',
        menu: 0
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
        },
        setName(state, action) {
            state.name = action.payload
        },
        setMenu(state, action) {
            state.menu = action.payload
        }
    }
})

export const { setAuth, setMenu, setName, setPosition, setRole } = positionSlice.actions

export default positionSlice.reducer