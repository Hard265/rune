lexer grammar RuneLexer;

INT              : 'int';
STR                  : 'str';
BOOL             : 'bool';
DOUBLE           : 'double';

STRING_LITERAL   : '"' .*? '"';
DOUBLE_LITERAL   : [0-9]+ '.' [0-9]+;
INTEGER_LITERAL  : [0-9]+;
BOOLEAN_LITERAL  : 'true' | 'false';

IDENTIFIER       : [a-zA-Z_][a-zA-Z0-9_]*;

NEWLINE          : '\r'? '\n' -> skip;
WHITESPACE       : [ \t]+ -> skip;

LPAR             : '(';
LSQB             : '[';
LBRACE           : '{';
RPAR             : ')';
RSQB             : ']';
RBRACE           : '}';
COLON            : ':';
COMMA            : ',';
SEMI             : ';';
PLUS             : '+';
MINUS            : '-';
STAR             : '*';
SLASH            : '/';
VBAR             : '|';
AMPER            : '&';
BANG             : '!';
QUESTION         : '?';
LESS             : '<';
GREATER          : '>';
EQUAL            : '=';
DOT              : '.';
PERCENT          : '%';
BACKQUOTE        : '`';
EQEQUAL          : '==';
NOTEQUAL         : '!=';
LESSEQUAL        : '<=';
GREATEREQUAL     : '>=';
TILDE            : '~';
CIRCUMFLEX       : '^';
LEFTSHIFT        : '<<';
RIGHTSHIFT       : '>>';
DOUBLEVBAR       : '||';
DOUBLEAMPER      : '&&';
DOUBLESTAR       : '**';
PLUSEQUAL        : '+=';
MINEQUAL         : '-=';
STAREQUAL        : '*=';
SLASHEQUAL       : '/=';
PERCENTEQUAL     : '%=';
AMPEREQUAL       : '&=';
VBAREQUAL        : '|=';
CIRCUMFLEXEQUAL  : '^=';
AT               : '@';

COMMENT          : '#?' ~[\r\n]* -> skip;
