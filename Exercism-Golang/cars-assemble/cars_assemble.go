package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	var success float64
    
    if speed ==0 {
        success = 0
    }
	
    if speed >= 1 && speed <= 4 {
    	success = 1
    }
	
    if speed >= 5 && speed <= 8 {
    	success = 0.9
    }
	
    if speed >=9 && speed <= 10 {
    	success = 0.77
    }

    return success
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {

    var success float64
    
    if speed ==0 {
        success = 0
    }
	
    if speed >= 1 && speed <= 4 {
    	success = 1
    }
	
    if speed >= 5 && speed <= 8 {
    	success = 0.9
    }
	
    if speed >=9 && speed <= 10 {
    	success = 0.77
    }

	return float64(float64(speed) * 221 * success)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {

    var success float64
    
    if speed ==0 {
        success = 0
    }
	
    if speed >= 1 && speed <= 4 {
    	success = 1
    }
	
    if speed >= 5 && speed <= 8 {
    	success = 0.9
    }
	
    if speed >=9 && speed <= 10 {
    	success = 0.77
    }

    result := (float64(speed) * 221 * success)/60
	return int(result)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	var produced float64
    produced = float64(speed * 221)
    if produced > limit {
        produced = limit
    }
	return produced
}
