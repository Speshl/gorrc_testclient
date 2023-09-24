package command

// import (
// 	"fmt"
// 	"log"

// 	"github.com/Speshl/gorrc_client/internal/config"
// 	"github.com/googolgl/go-i2c"
// 	"github.com/googolgl/go-pca9685"
// )

// const (
// 	MaxValue = 1.0
// 	MinValue = -1.0
// 	MaxPulse = pca9685.ServoMaxPulseDef
// 	MinPulse = pca9685.ServoMinPulseDef
// 	AcRange  = pca9685.ServoRangeDef

// 	MaxSupportedServos = 16
// )

// type Command struct {
// 	cfg    config.CommandConfig
// 	servos map[string]*Servo
// 	driver *pca9685.PCA9685
// }

// type Servo struct {
// 	name     string
// 	inverted bool
// 	offset   float64
// 	servo    *pca9685.Servo
// }

// func NewTestCommand(cfg config.CommandConfig) *Command {
// 	return &Command{
// 		cfg:    cfg,
// 		servos: make(map[string]*Servo, MaxSupportedServos),
// 	}
// }

// func (c *Command) Init() error {
// 	i2c, err := i2c.New(c.cfg.Address, c.cfg.I2CDevice)
// 	if err != nil {
// 		return fmt.Errorf("error starting i2c with address - %w", err)
// 	}

// 	c.driver, err = pca9685.New(i2c, nil)
// 	if err != nil {
// 		return fmt.Errorf("error getting servo driver - %w", err)
// 	}

// 	for i := range c.cfg.ServoCfgs {
// 		name := c.cfg.ServoCfgs[i].Name
// 		c.servos[name].name = name
// 		c.servos[name].inverted = c.cfg.ServoCfgs[i].Inverted
// 		c.servos[name].offset = float64(c.cfg.ServoCfgs[i].Offset) / 100
// 		c.servos[name].servo = c.driver.ServoNew(c.cfg.ServoCfgs[i].Channel, &pca9685.ServOptions{
// 			AcRange:  AcRange,
// 			MinPulse: float32(c.cfg.ServoCfgs[i].MinPulse),
// 			MaxPulse: float32(c.cfg.ServoCfgs[i].MaxPulse),
// 		})
// 		log.Printf("%s servo added\n", name)
// 	}
// 	return nil
// }

// func (c *Command) Set(name string, value, min, max float64) error {
// 	val, ok := c.servos[name]
// 	if ok {
// 		mappedValue := mapToRange(value+val.offset, min, max, MinValue, MaxValue)
// 		if c.servos[name].inverted {
// 			mappedValue = MaxValue - mappedValue
// 		}

// 		err := c.servos[name].servo.Fraction(float32(mappedValue))
// 		if err != nil {
// 			return fmt.Errorf("failed setting servo value - name: %s value: %d - error: %w\n", mappedValue, err)
// 		}
// 		log.Printf("Servo %s: value: %.2f offset: %.2f mapped: %.2f\n", name, value, val.offset, mappedValue)
// 	}
// 	return nil
// }

// func mapToRange(value, min, max, minReturn, maxReturn float64) float64 {
// 	mappedValue := (maxReturn-minReturn)*(value-min)/(max-min) + minReturn

// 	if mappedValue > maxReturn {
// 		return maxReturn
// 	} else if mappedValue < minReturn {
// 		return minReturn
// 	} else {
// 		return mappedValue
// 	}
// }
