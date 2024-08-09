import { createSlice } from '@reduxjs/toolkit'

export const positionSlice = createSlice({
    name: 'position',
    initialState: {
        position: 'login',
        auth: ''
    },
    reducers: {
        setPosition(state, action) {
            state.position = action.payload
        },
        setAuth(state, action) {
            state.auth = action.payload
        }
    }
})

export const { setAuth, setPosition } = positionSlice.actions

export default positionSlice.reducer