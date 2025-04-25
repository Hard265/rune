grammar RuneParser;

options {
    tokenVocab = RuneLexer;
}

program: statement* EOF;

statement
    : declaration
    | assignment
    | declarationAssignment
    ;

declaration: type_ IDENTIFIER;

assignment: IDENTIFIER EQUAL expression;

declarationAssignment: type_ IDENTIFIER EQUAL expression;

type_: INT | BOOL | STR | DOUBLE;

expression
    : unaryExpression
    ;

unaryExpression
    : (PLUS | MINUS | BANG) unaryExpression
    | binaryExpression
    ;

binaryExpression
    : logicalOrExpression
    ;

logicalOrExpression
    : logicalAndExpression (DOUBLEVBAR logicalAndExpression)*
    ;

logicalAndExpression
    : equalityExpression (DOUBLEAMPER equalityExpression)*
    ;

equalityExpression
    : relationalExpression ((EQEQUAL | NOTEQUAL) relationalExpression)*
    ;

relationalExpression
    : additiveExpression ((LESS | LESSEQUAL | GREATER | GREATEREQUAL) additiveExpression)*
    ;

additiveExpression
    : multiplicativeExpression ((PLUS | MINUS) multiplicativeExpression)*
    ;

multiplicativeExpression
    : primaryExpression ((STAR | SLASH) primaryExpression)*
    ;

primaryExpression
    : literal
    | IDENTIFIER
    | LPAR expression RPAR
    ;

literal
    : INTEGER_LITERAL
    | BOOLEAN_LITERAL
    | STRING_LITERAL
	| DOUBLE_LITERAL
    ;

