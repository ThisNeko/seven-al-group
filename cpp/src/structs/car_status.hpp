#ifndef CAR_STATUS_H
#define CAR_STATUS_H

typedef struct {
    int X;
    int Y;
} Position;

typedef struct {
    int X = 0;
    int Y = 0;
} Vitesse;

typedef struct {
    int X = 0;
    int Y = 0;
} Acceleration;

typedef struct {
    Position position;
    Vitesse vitesse;
    bool panne;
    int ID;
} CarStatus;

#endif // CAR_STATUS_H