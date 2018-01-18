#ifndef CAR_STATUS_H
#define CAR_STATUS_H

typedef struct {
    int X;
    int Y;
} Position;

typedef struct {
    int X = 0;
    int Y = 0;
} Speed;

typedef struct {
    int X = 0;
    int Y = 0;
} Acceleration;

typedef struct {
    Position position;
    Speed speed;
    Acceleration acceleration;
    bool brakes;
    int ID;
} CarStatus;

#endif // CAR_STATUS_H