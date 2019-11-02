package main
 
import (
    "logger"
)
 
func main() {
    logger.Infof("info simple zap logger example")    
    logger.Debugf("debug simple zap logger example")    
    logger.Errorf("error simple zap logger example","test for another args")    
}
