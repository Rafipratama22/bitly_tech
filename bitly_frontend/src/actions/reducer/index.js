import { combineReducers } from 'redux'
import User from './Reducer'

const allReducer = combineReducers({
    user: User
})

export default allReducer